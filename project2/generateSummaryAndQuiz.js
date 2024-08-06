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
                { role: 'user', content: ` סכם את הטקסט הבא בפירוט ככל האפשר והשתמש בכמה מילים שאתה יכול (כמה שיותר). הסבר את הנושאים שהוסברו בסרטון תוך נתינת דוגמאות\n\n${text}` }
            ],
            max_tokens: 1000,
        });

        
        const quizResponse = await openai.chat.completions.create({
            model: 'gpt-3.5-turbo',
            messages: [
                { role: 'system', content: 'You are a helpful assistant.' },
                { role: 'user', content: `צור חידון אמריקאי-כלומר שלכל שאלה יש 4 אופציות לתשובה. תכתוב תחילה את כל השאלות ואת ארבעת האופציות לפיתרון של כל אחד, ככה שמתחת לכל שאלה תופיע 4 האופציות שלה. רק לאחר מכן תכתוב איזו תשובה היא נכונה לכל אחת מהשאלות. תכתוב שאלות שנוגעות רק לחומר הלימודי של הטקסט ולא לפרטים יבשים כמו שם המרצה או איפה הוא למד, תיעזר בחומר מהאינטרנט שקשור לנושא המלומד ואל תסתמך רק על הטקסט. :\n\n${text}` }
            ],
            max_tokens: 1000,
        });

        const summary = summaryResponse.choices[0].message.content.trim();
        const quiz = quizResponse.choices[0].message.content.trim();

        fs.writeFileSync('summary.txt', summary, 'utf8');
        fs.writeFileSync('quiz.txt', quiz, 'utf8');


    } catch (error) {
        console.error('Error generating summary and quiz:', error);
    }
}

module.exports = generateSummaryAndQuiz;

(async () => {
    const text = "הכנס כאן את הטקסט שלך לסיכום ולחידון";
    await generateSummaryAndQuiz(text);
})();
