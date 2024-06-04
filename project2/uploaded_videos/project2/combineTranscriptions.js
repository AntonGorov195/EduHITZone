const fs = require('fs');
const path = require('path');

function combineTranscriptions(directoryPath, outputPath) {
    return new Promise((resolve, reject) => {
        fs.readdir(directoryPath, (err, files) => {
            if (err) {
                return reject(err);
            }

            const transcriptionFiles = files.filter(file => path.extname(file) === '.txt');
            transcriptionFiles.sort((a, b) => a.localeCompare(b, undefined, { numeric: true }));

            let combinedTranscription = '';
            transcriptionFiles.forEach(file => {
                const filePath = path.join(directoryPath, file);
                const content = fs.readFileSync(filePath, 'utf-8');
                combinedTranscription += content + ' ';
            });

            fs.writeFileSync(outputPath, combinedTranscription.trim());
            resolve(outputPath);
        });
    });
}

module.exports = combineTranscriptions;
