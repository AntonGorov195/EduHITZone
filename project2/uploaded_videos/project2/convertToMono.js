const ffmpeg = require('fluent-ffmpeg');
const path = require('path');

function convertToMono(inputAudioPath) {
    return new Promise((resolve, reject) => {
        const outputMonoAudioPath = path.join(path.dirname(inputAudioPath), 'output_audio_mono.mp3');
        ffmpeg(inputAudioPath)
            .audioChannels(1)
            .output(outputMonoAudioPath)
            .on('end', () => {
                console.log('Audio conversion to mono completed.');
                resolve(outputMonoAudioPath);
            })
            .on('error', (err) => {
                console.log('Error converting audio to mono:', err);
                reject(err);
            })
            .run();
    });
}

module.exports = convertToMono;
