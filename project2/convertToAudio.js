const ffmpeg = require('fluent-ffmpeg');
const path = require('path');

function convertToAudio(videoPath) {
    return new Promise((resolve, reject) => {
        const outputAudioPath = path.join(path.dirname(videoPath), 'output_audio.mp3');
        ffmpeg(videoPath)
            .output(outputAudioPath)
            .on('end', () => {
                console.log('Audio conversion completed.');
                resolve(outputAudioPath);
            })
            .on('error', (err) => {
                console.log('Error converting video to audio:', err);
                reject(err);
            })
            .run();
    });
}

module.exports = convertToAudio;
