// const { OpenAIApi } = require('openai');
// require('dotenv').config();

// const openai = new OpenAIApi({
//   apiKey: process.env.OPENAI_API_KEY,
// });

// async function testAPI() {
//     try {
//         const response = await openai.createChatCompletion({
//             model: 'gpt-3.5-turbo',
//             messages: [
//                 { role: 'system', content: 'You are a helpful assistant.' },
//                 { role: 'user', content: 'Say hello!' }
//             ],
//             max_tokens: 10,
//         });
//         console.log(response.data.choices[0].message.content.trim());
//     } catch (error) {
//         console.error('Error with OpenAI API:', error);
//     }
// }

// testAPI();


// runTest();


const OpenAI = require('openai');
require('dotenv').config();

const openai = new OpenAI({
    apiKey: process.env.OPENAI_API_KEY,
});

async function testAPI() {
    try {
        const response = await openai.chat.completions.create({
            model: 'gpt-3.5-turbo',
            messages: [
                { role: 'system', content: 'You are a helpful assistant.' },
                { role: 'user', content: 'Say hello!' }
            ],
            max_tokens: 10,
        });
        console.log(response.choices[0].message.content.trim());
    } catch (error) {
        console.error('Error with OpenAI API:', error);
    }
}


testAPI();



