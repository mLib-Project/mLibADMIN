const axios = require('axios')
const books = require('./books')
const code = require('./code.env')

const pushBook = async (book) => {
    try {
        const res = await axios.post('http://localhost/api', book)
    } catch (err) {
      console.log(err)
    }
}

const dropBooks = async () => {
    try {
        await axios.delete('http://localhost/api/' + code)
    } catch (err) {
      console.log(err)
    }
}

const pushData = () => {
    for(let i=0; i <= books.array.length - 1; i++ ) {
        pushBook(books.array[i])
    }
}

dropBooks()
pushData()
