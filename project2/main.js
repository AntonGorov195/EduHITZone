const path = require('path');
const fs = require('fs');
const installDependencies = require('./installDependencies');
const uploadVideo = require('./uploadVideo');
const convertToAudio = require('./convertToAudio');
const convertAudioToMono = require('./convertAudioToMono');
const transcribeAudio = require('./transcribeAudio');
const generateSummaryAndQuiz = require('./generateSummaryAndQuiz')

async function main() {
    try {
        await installDependencies();

        const videoPath = await uploadVideo();
        //console.log(`Video uploaded successfully to: ${videoPath}`);

        const audioPath = await convertToAudio(videoPath);
        //console.log(`Audio conversion completed.`);

        const monoAudioPath = await convertAudioToMono(audioPath);
        //console.log(`Audio conversion to mono completed.`);

        const transcription = await transcribeAudio(monoAudioPath);
        //console.log(`Transcription completed: ${transcription}`);

        await generateSummaryAndQuiz(transcription);
        console.log(`Summary and quiz generated successfuly.`);

    } catch (error) {
        console.error('Error during processing:', error);
    }
}

main();
