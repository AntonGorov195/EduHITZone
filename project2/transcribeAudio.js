const fs = require('fs');
const path = require('path');
const { Storage } = require('@google-cloud/storage');
const speech = require('@google-cloud/speech');
const ffmpeg = require('fluent-ffmpeg');

const client = new speech.SpeechClient({
    keyFilename: 'my-project-423811-ad4a844d02cf.json'
});

const storage = new Storage({
    keyFilename: 'my-project-423811-ad4a844d02cf.json'
});

const bucketName = 'my-project-bucket2';

async function transcribeAudio(audioFilePath) {
    try {
        const convertedAudioPath = audioFilePath.replace('.mp3', '_converted.wav');
        await convertToWav(audioFilePath, convertedAudioPath);

        await uploadToBucket(convertedAudioPath);

        const gcsUri = `gs://${bucketName}/${path.basename(convertedAudioPath)}`;
        const transcription = await transcribeLongAudio(gcsUri);

        const outputFilePath = path.join(__dirname, 'combined_transcription.txt');
        fs.writeFileSync(outputFilePath, transcription, 'utf8');

        return transcription;

    } catch (error) {
        console.error(`Error transcribing ${audioFilePath}:`, error);
        throw error;
    }
}

async function convertToWav(inputPath, outputPath) {
    return new Promise((resolve, reject) => {
        ffmpeg(inputPath)
            .outputOptions([
                '-ar 16000',
                '-ac 1'
            ])
            .output(outputPath)
            .on('end', () => {
                resolve(outputPath);
            })
            .on('error', reject)
            .run();
    });
}

async function uploadToBucket(filePath) {
    await storage.bucket(bucketName).upload(filePath, {
        destination: path.basename(filePath),
    });
}

async function transcribeLongAudio(gcsUri) {
    const request = {
        audio: {
            uri: gcsUri,
        },
        config: {
            encoding: 'LINEAR16',
            sampleRateHertz: 16000,
            languageCode: 'he-IL',
        },
    };

    const [operation] = await client.longRunningRecognize(request);
    const [response] = await operation.promise();

    if (!response.results || response.results.length === 0) {
        console.error(`No transcription results for ${gcsUri}`);
        return '';
    }

    const transcription = response.results
        .map(result => result.alternatives[0].transcript)
        .join('\n');
    return transcription;
}

module.exports = transcribeAudio;

