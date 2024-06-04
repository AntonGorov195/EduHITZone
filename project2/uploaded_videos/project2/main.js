const path = require('path');
const fs = require('fs');
const installDependencies = require('./installDependencies');
const uploadVideo = require('./uploadVideo');
const convertToAudio = require('./convertToAudio');
const convertAudioToMono = require('./convertAudioToMono');
const splitAudio = require('./splitAudio');
const transcribeAudio = require('./transcribeAudio');

async function main() {
    try {
        await installDependencies();

        const videoPath = await uploadVideo();
        console.log(`Video uploaded successfully to: ${videoPath}`);

        const audioPath = await convertToAudio(videoPath);
        console.log(`Audio conversion completed.`);

        const monoAudioPath = await convertAudioToMono(audioPath);
        console.log(`Audio conversion to mono completed.`);

        const audioChunks = await splitAudio(monoAudioPath);
        console.log(`Audio chunks created:`, audioChunks);

        const transcriptions = [];
        for (const chunk of audioChunks) {
            const transcription = await transcribeAudio(chunk);
            transcriptions.push(transcription);
        }

        const combinedTranscription = transcriptions.join(' ');
        const combinedTranscriptionPath = path.join(__dirname, 'combined_transcription.txt');
        fs.writeFileSync(combinedTranscriptionPath, combinedTranscription, 'utf8');
        console.log(`Combined transcription saved to: ${combinedTranscriptionPath}`);

    } catch (error) {
        console.error('Error during processing:', error);
    }
}

main();
