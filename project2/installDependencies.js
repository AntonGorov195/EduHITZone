const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

async function installDependencies() {
    const dependencies = [
        { name: '@google-cloud/speech', version: null },
        { name: '@google-cloud/storage', version: null },
        { name: 'fluent-ffmpeg', version: null },
        { name: 'ffmpeg-static', version: null },
        { name: 'openai', version: '^4.0.0' },
        { name: 'dotenv', version: null }
    ];

    for (const dep of dependencies) {
        try {
            const packagePath = require.resolve(path.join(dep.name, 'package.json'));
            const packageJson = JSON.parse(fs.readFileSync(packagePath, 'utf8'));
            if (dep.version && packageJson.version !== dep.version.replace('^', '')) {
                throw new Error(`Version mismatch for ${dep.name}`);
            }
            console.log(`${dep.name} is already installed.`);
        } catch (e) {
            const packageToInstall = dep.version ? `${dep.name}@${dep.version}` : dep.name;
            console.log(`Installing package: ${packageToInstall}`);
            execSync(`npm install ${packageToInstall}`, { stdio: 'inherit' });
            console.log(`${dep.name} installed successfully.`);
        }
    }
}

module.exports = installDependencies;

