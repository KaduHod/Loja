import knex from 'knex'
import * as config from 'dotenv';
config.config()
export const db = knex({
    client: "pg",
    connection: {
        host: process.env.POSTGRE_HOST,
        port: process.env.POSTGREE_PORTA,
        user: process.env.POSTGRE_USER ,
        password: process.env.POSTGRE_PWD ,
        database: process.env.POSTGRE_DB
    },
})
