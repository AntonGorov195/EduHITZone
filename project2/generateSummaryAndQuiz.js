const fs = require('fs');
const OpenAI = require('openai');
require('dotenv').config();

const openai = new OpenAI({
    apiKey: process.env.OPENAI_API_KEY,
});

async function generateSummaryAndQuiz(text) {
    try {
        // יצירת סיכום
        const summaryResponse = await openai.chat.completions.create({
            model: 'gpt-3.5-turbo',
            messages: [
                { role: 'system', content: 'You are a helpful assistant.' },
                { role: 'user', content: `סכם את הטקסט הבא:\n\n${text}` }
            ],
            max_tokens: 150,
        });

        // יצירת חידון
        const quizResponse = await openai.chat.completions.create({
            model: 'gpt-3.5-turbo',
            messages: [
                { role: 'system', content: 'You are a helpful assistant.' },
                { role: 'user', content: `צור חידון אמריקאי מתוך הטקסט הבא:\n\n${text}` }
            ],
            max_tokens: 150,
        });

        const summary = summaryResponse.choices[0].message.content.trim();
        const quiz = quizResponse.choices[0].message.content.trim();

        fs.writeFileSync('summary.txt', summary, 'utf8');
        fs.writeFileSync('quiz.txt', quiz, 'utf8');

        console.log('Summary and quiz have been generated and saved.');
    } catch (error) {
        console.error('Error generating summary and quiz:', error);
    }
}

module.exports = generateSummaryAndQuiz;

(async () => {
    const text = "הכנס כאן את הטקסט שלך לסיכום ולחידון";
    await generateSummaryAndQuiz(text);
})();




// const axios = require('axios');

// // הגדרת מפתח ה-API שלך
// const apiKey = '9fa2466043424b00b238aa3ca7b57da8.68a83cbbf18f2d65'; 
// const apiUrl = 'https://api.copilot.com';  

// async function generateSummaryAndQuiz(text) {
//     try {
//         // בקשה ליצירת סיכום
//         const summaryResponse = await axios.post(`${apiUrl}/summarize`, {
//             text: text,
//             language: 'he'
//         }, {
//             headers: {
//                 'Authorization': `Bearer ${apiKey}`
//             }
//         });
//         const summary = summaryResponse.data.summary;

//         const quizResponse = await axios.post(`${apiUrl}/quiz`, {
//             text: text,
//             language: 'he',
//             numQuestions: 5
//         }, {
//             headers: {
//                 'Authorization': `Bearer ${apiKey}`
//             }
//         });
//         const quiz = quizResponse.data.quiz;

//         return { summary, quiz };
//     } catch (error) {
//         console.error('Error generating summary and quiz:', error);
//         throw error;
//     }
// }

// const videoText = 'הטקסט שהתקבל מהוידאו לאחר המרת אודיו לטקסט'; 
// generateSummaryAndQuiz(videoText).then(({ summary, quiz }) => {
//     console.log('Summary:', summary);
//     console.log('Quiz:', quiz);
// }).catch(error => {
//     console.error('Error:', error);
// });
// module.exports = generateSummaryAndQuiz