const ffmpeg = require('fluent-ffmpeg');
const path = require('path');

function convertAudioToMono(audioPath) {
    return new Promise((resolve, reject) => {
        const outputFilePath = path.join(path.dirname(audioPath), 'mono_' + path.basename(audioPath));
        ffmpeg(audioPath)
            .audioChannels(1)
            .on('end', () => {
                console.log('Audio conversion to mono completed.');
                resolve(outputFilePath);
            })
            .on('error', (err) => {
                console.error('Error converting audio to mono:', err);
                reject(err);
            })
            .save(outputFilePath);
    });
}

module.exports = convertAudioToMono;
