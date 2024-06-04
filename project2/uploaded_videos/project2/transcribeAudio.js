const fs = require('fs');
const speech = require('@google-cloud/speech');

// ודא שהמפתח מוגדר בצורה נכונה
const client = new speech.SpeechClient({
    keyFilename: 'C:/Users/TLV-9/Downloads/my-project-423811-ad4a844d02cf.json'
});

async function transcribeAudio(audioFilePath) {
    try {
        console.log(`Transcribing ${audioFilePath}...`);

        const file = fs.readFileSync(audioFilePath);
        const audioBytes = file.toString('base64');

        const request = {
            audio: {
                content: audioBytes,
            },
            config: {
                encoding: 'LINEAR16',
                sampleRateHertz: 16000,
                languageCode: 'he-IL', // הגדרת השפה לעברית
            },
        };

        console.log('Sending request to Google API with config:', request.config);

        const [response] = await client.recognize(request);

        if (response.results.length === 0) {
            console.error(`No transcription results for ${audioFilePath}`);
            return '';
        }

        const transcription = response.results
            .map(result => result.alternatives[0].transcript)
            .join('\n');

        console.log(`Transcription for ${audioFilePath}: ${transcription}`); // הדפסת התמלול למסך

        const outputFilePath = audioFilePath.replace('audio_chunks', 'transcriptions').replace('.wav', '.txt');
        fs.writeFileSync(outputFilePath, transcription, 'utf8'); // כתיבה לקובץ
        return transcription;

    } catch (error) {
        console.error(`Error transcribing ${audioFilePath}:`, error);
        throw error;
    }
}

module.exports = transcribeAudio;
