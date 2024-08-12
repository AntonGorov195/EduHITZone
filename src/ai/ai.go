package ai

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"cloud.google.com/go/storage"
	openai "github.com/sashabaranov/go-openai"
	"google.golang.org/api/option"
)

func uploadVideo(path string) (string, error) {
	upload_dir := "./uploaded_videos"
	err := os.MkdirAll(upload_dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	path_base := filepath.Base(path)
	dest_path := filepath.Join(upload_dir, path_base)

	in, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer in.Close()
	out, err := os.Create(dest_path)
	if err != nil {
		return "", err
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return "", err
	}
	err = out.Sync()
	if err != nil {
		return "", err
	}
	return dest_path, nil
}
func convertToAudio(video_path string) (string, error) {
	outputAudioPath := filepath.Join(filepath.Dir(video_path), "output_audio.mp3")
	cmd := exec.Command("ffmpeg", "-i", video_path, "-q:a", "0", "-map", "a", outputAudioPath)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error converting video to audio: %v\n", err)
		return "", err
	}

	fmt.Println("Audio conversion completed.")
	return outputAudioPath, nil
}
func convertAudioToMono(audioPath string) (string, error) {
	outputFilePath := filepath.Join(filepath.Dir(audioPath), "mono_"+filepath.Base(audioPath))

	cmd := exec.Command("ffmpeg", "-i", audioPath, "-ac", "1", outputFilePath)

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error converting audio to mono: %v\n", err)
		return "", err
	}

	fmt.Println("Audio conversion to mono completed.")
	return outputFilePath, nil
}
func generateSummaryAndQuiz(text string) (summary, quiz string, err error) {
	client := openai.NewClient("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") // Replace with your OpenAI API key

	fmt.Println(text)
	// Generate Summary
	summaryRequest := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: fmt.Sprintf(`סכם את הטקסט הבא בפירוט ככל האפשר והשתמש בכמה מילים שאתה יכול (כמה שיותר). הסבר את הנושאים שהוסברו בסרטון תוך נתינת דוגמאות\n\n%s`, text)},
		},
		MaxTokens: 1000,
	}

	summaryResponse, err := client.CreateChatCompletion(context.Background(), summaryRequest)
	if err != nil {
		return "", "", fmt.Errorf("error generating summary: %v", err)
	}
	summary = summaryResponse.Choices[0].Message.Content
	// Write Summary to File
	err = os.WriteFile("summary.txt", []byte(summary), 0644)
	if err != nil {
		return "", "", err
	}
	// Generate Quiz
	quizRequest := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: fmt.Sprintf(`צור חידון אמריקאי-כלומר שלכל שאלה יש 4 אופציות לתשובה. תכתוב תחילה את כל השאלות ואת ארבעת האופציות לפיתרון של כל אחד, ככה שמתחת לכל שאלה תופיע 4 האופציות שלה. רק לאחר מכן תכתוב איזו תשובה היא נכונה לכל אחת מהשאלות. תכתוב שאלות שנוגעות רק לחומר הלימודי של הטקסט ולא לפרטים יבשים כמו שם המרצה או איפה הוא למד, תיעזר בחומר מהאינטרנט שקשור לנושא המלומד ואל תסתמך רק על הטקסט. :\n\n%s`, text)},
		},
		MaxTokens: 1000,
	}

	quizResponse, err := client.CreateChatCompletion(context.Background(), quizRequest)
	if err != nil {
		return "", "", fmt.Errorf("error generating quiz: %v", err)
	}
	quiz = quizResponse.Choices[0].Message.Content
	// Write Quiz to File
	err = os.WriteFile("quiz.txt", []byte(quiz), 0644)
	if err != nil {
		return "", "", err
	}
	return summary, quiz, nil
}
func GenerateAIContent(video string) (summary, quiz string, err error) {
	fmt.Println("Begin generating ai content ")
	fmt.Println(video)
	// path, err := uploadVideo(video)
	// if err != nil {
	// 	return "", "", err
	// }
	path, err := convertToAudio(video)
	if err != nil {
		return "", "", err
	}
	path, err = convertAudioToMono(path)
	if err != nil {
		return "", "", err
	}
	text, err := transcribeAudio(path)
	if err != nil {
		return "", "", err
	}
	fmt.Println("AI finished")
	// return text, "", err
	return generateSummaryAndQuiz(text)
}

var (
	bucketName    = "my-project-bucket2"
	keyFilename   = "google-cred.json.json"
	outputFileDir = "./" // Change this to your desired output directory
)

func transcribeAudio(audioFilePath string) (string, error) {
	convertedAudioPath := strings.Replace(audioFilePath, ".mp3", "_converted.wav", 1)

	err := convertToWav(audioFilePath, convertedAudioPath)
	if err != nil {
		return "", fmt.Errorf("error converting audio: %v", err)
	}

	err = uploadToBucket(convertedAudioPath)
	if err != nil {
		return "", fmt.Errorf("error uploading to bucket: %v", err)
	}

	gcsUri := fmt.Sprintf("gs://%s/%s", bucketName, filepath.Base(convertedAudioPath))
	transcription, err := transcribeLongAudio(gcsUri)
	if err != nil {
		return "", fmt.Errorf("error transcribing audio: %v", err)
	}

	outputFilePath := filepath.Join(outputFileDir, "combined_transcription.txt")
	err = ioutil.WriteFile(outputFilePath, []byte(transcription), 0644)
	if err != nil {
		return "", fmt.Errorf("error writing transcription to file: %v", err)
	}
	fmt.Println("Transcription finished")

	return transcription, nil
}

func convertToWav(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ar", "16000", "-ac", "1", outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("ffmpeg conversion failed: %v", err)
	}
	return nil
}

func uploadToBucket(filePath string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(keyFilename))
	if err != nil {
		return fmt.Errorf("failed to create storage client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	object := bucket.Object(filepath.Base(filePath))
	wc := object.NewWriter(ctx)
	defer wc.Close()

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	if _, err = wc.Write(data); err != nil {
		return fmt.Errorf("failed to write file to bucket: %v", err)
	}

	return nil
}

func transcribeLongAudio(gcsUri string) (string, error) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx, option.WithCredentialsFile(keyFilename))
	if err != nil {
		return "", fmt.Errorf("failed to create speech client: %v", err)
	}
	defer client.Close()

	req := &speechpb.LongRunningRecognizeRequest{
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Uri{Uri: gcsUri},
		},
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_LINEAR16,
			SampleRateHertz: 16000,
			LanguageCode:    "he-IL",
		},
	}

	op, err := client.LongRunningRecognize(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to start long-running recognize: %v", err)
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for long-running recognize response: %v", err)
	}

	if len(resp.Results) == 0 {
		log.Printf("No transcription results for %s", gcsUri)
		return "", nil
	}

	var transcription string
	for _, result := range resp.Results {
		transcription += result.Alternatives[0].Transcript + "\n"
	}

	return transcription, nil
}
