const fs = require('fs');
const ffmpeg = require('fluent-ffmpeg');
const path = require('path');

function splitAudio(inputPath) {
    return new Promise((resolve, reject) => {
        const outputDir = path.join(__dirname, 'audio_chunks');
        if (!fs.existsSync(outputDir)) {
            fs.mkdirSync(outputDir);
        }

        // Delete old chunks if exist
        fs.readdirSync(outputDir).forEach(file => {
            fs.unlinkSync(path.join(outputDir, file));
        });

        ffmpeg(inputPath)
            .output(path.join(outputDir, 'chunk_%03d.wav'))
            .outputOptions([
                '-f segment',
                '-segment_time 60', // Split every 60 seconds
                '-c copy'
            ])
            .on('end', function () {
                const audioChunks = fs.readdirSync(outputDir)
                    .filter(file => file.endsWith('.wav'))
                    .map(file => path.join(outputDir, file));
                
                console.log('Audio chunks created:', audioChunks);
                resolve(audioChunks);
            })
            .on('error', function (err) {
                reject(err);
            })
            .run();
    });
}

module.exports = splitAudio;
