const { OpenAIApi } = require("openai");
require('dotenv').config();

const openai = new OpenAIApi({
  apiKey: process.env.OPENAI_API_KEY,
});

async function runTest() {
  try {
    const response = await openai.createCompletion({
      model: "text-davinci-003",
      prompt: "Say this is a test",
      max_tokens: 5,
    });
    console.log(response.data.choices[0].text);
  } catch (error) {
    console.error('Error with OpenAI API:', error);
  }
}

runTest();

