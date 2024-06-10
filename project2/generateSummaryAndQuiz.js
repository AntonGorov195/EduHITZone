const fs = require('fs');
const { Configuration, OpenAIApi } = require('openai');
require('dotenv').config();

const configuration = new Configuration({
    apiKey: process.env.OPENAI_API_KEY,
});
const openai = new OpenAIApi(configuration);

async function generateSummaryAndQuiz(text) {
    try {
        const summaryResponse = await openai.createChatCompletion({
            model: 'gpt-3.5-turbo',
            messages: [
                { role: 'system', content: 'You are a helpful assistant.' },
                { role: 'user', content: `סכם את הטקסט הבא:\n\n${text}` }
            ],
            max_tokens: 150,
        });

        const quizResponse = await openai.createChatCompletion({
            model: 'gpt-3.5-turbo',
            messages: [
                { role: 'system', content: 'You are a helpful assistant.' },
                { role: 'user', content: `צור חידון אמריקאי מתוך הטקסט הבא:\n\n${text}` }
            ],
            max_tokens: 150,
        });

        const summary = summaryResponse.data.choices[0].message.content.trim();
        const quiz = quizResponse.data.choices[0].message.content.trim();

        fs.writeFileSync('summary.txt', summary, 'utf8');
        fs.writeFileSync('quiz.txt', quiz, 'utf8');

        console.log('Summary and quiz have been generated and saved.');
    } catch (error) {
        console.error('Error generating summary and quiz:', error);
    }
}

module.exports = generateSummaryAndQuiz;


