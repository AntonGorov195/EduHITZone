const { execSync } = require('child_process');

async function installDependencies() {
    const dependencies = [
        '@google-cloud/speech',
        'fluent-ffmpeg',
        'ffmpeg-static'
    ];

    for (const dep of dependencies) {
        console.log(`Installing package: ${dep}`);
        execSync(`npm install ${dep}`, { stdio: 'inherit' });
        console.log(`${dep} installed successfully.`);
    }
}

module.exports = installDependencies;
