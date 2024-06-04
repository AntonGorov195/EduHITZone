const { execSync } = require('child_process');

async function installDependencies() {
    const dependencies = [
        '@google-cloud/speech',
        '@google-cloud/storage',
        'fluent-ffmpeg',
        'ffmpeg-static'
    ];

    for (const dep of dependencies) {
        try {
            require.resolve(dep);
            console.log(`${dep} is already installed.`);
        } catch (e) {
            console.log(`Installing package: ${dep}`);
            execSync(`npm install ${dep}`, { stdio: 'inherit' });
            console.log(`${dep} installed successfully.`);
        }
    }
}

module.exports = installDependencies;
