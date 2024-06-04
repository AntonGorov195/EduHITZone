const convertToAudio = require('./convertToAudio');

const videoPath = './uploaded_videos/example_video_project.mp4';
const outputAudioPath = './uploaded_videos/output_audio.mp3';

convertToAudio(videoPath, outputAudioPath)
    .then(() => {
        console.log('Audio conversion test completed successfully.');
    })
    .catch((error) => {
        console.error('Error in audio conversion test:', error);
    });
