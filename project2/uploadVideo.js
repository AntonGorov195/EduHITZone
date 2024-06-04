const fs = require('fs');
const path = require('path');
const readline = require('readline');

function uploadVideo() {
    return new Promise((resolve, reject) => {
        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout
        });

        rl.question('Enter the path to the video file: ', (inputPath) => {
            // Remove quotes if present
            const videoPath = inputPath.replace(/['"]+/g, '');
            
            const uploadDir = './uploaded_videos';
            if (!fs.existsSync(uploadDir)) {
                fs.mkdirSync(uploadDir);
            } else {
                console.log(`Directory '${uploadDir}' already exists.`);
            }

            const fileName = path.basename(videoPath);
            const destinationPath = path.join(uploadDir, fileName);

            try {
                fs.copyFileSync(videoPath, destinationPath);
                console.log(`Video uploaded successfully to: ${destinationPath}`);
                rl.close();
                resolve(destinationPath);
            } catch (error) {
                console.error('Error uploading video:', error);
                rl.close();
                reject(error);
            }
        });
    });
}

module.exports = uploadVideo;
