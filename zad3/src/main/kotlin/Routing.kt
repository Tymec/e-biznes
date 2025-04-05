package com.example

import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import io.ktor.client.statement.*
import io.ktor.http.*
import io.ktor.client.*
import io.ktor.client.call.*

fun Application.configureRouting() {
    routing {
        post("/send") {
            val msg = call.receive<String>()
            val res = sendDiscordMessage(msg)

            if (res == null) {
                call.respondText("Missing environment variables", status = io.ktor.http.HttpStatusCode.InternalServerError)
            } else if (res.status.value in 200..299) {
                call.respondText("Message sent successfully")
            } else {
                call.respondText("Failed to send message\n${res.body<String>()}", status = io.ktor.http.HttpStatusCode.InternalServerError)
            }
        }
    }
}