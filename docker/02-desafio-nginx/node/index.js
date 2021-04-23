const express = require('express')
const mysql = require('mysql')

const app = express()
const port = 3000
const config = {
    host: 'db-gg2',
    user: 'root',
    password: 'root',
    database: 'gg2db'
}

app.get('/', (req, res) => {
    const connection = mysql.createConnection(config)
    connection.query(`create table if not exists people(id int not null auto_increment, name varchar(255), primary key(id))`)

    const ts = Date.now()
    const name = `Pessoa ${Math.floor(ts / 1000)}`
    connection.query(`INSERT INTO people(name) values ('${name}')`)

    connection.query(`SELECT * FROM people`, (err, rows) => {
        if (err) throw err
        let body = '<h1>Full Cycle Rocks!</h1>'
        body += '<ul>'
        rows.forEach(row => {
            body += `<li>${row.name}</li>`
        })
        body += '</ul>'
        res.send(body)
    })
    connection.end()
})

app.listen(port, () => {
    console.log(`Rodando na porta ${port}`)
})
