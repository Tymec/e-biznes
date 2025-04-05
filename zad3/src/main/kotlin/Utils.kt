package com.example

import io.ktor.client.*
import io.ktor.client.call.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.ktor.http.*

const val BASE_URL = "https://discord.com/api/webhooks/"

suspend fun sendDiscordMessage(message: String): HttpResponse? {
    val WEBHOOK_ID = System.getenv("WEBHOOK_ID") ?: ""
    val WEBHOOK_TOKEN = System.getenv("WEBHOOK_TOKEN") ?: ""

    if (WEBHOOK_ID.isEmpty() || WEBHOOK_TOKEN.isEmpty()) {
        return null
    }

    val url = "$BASE_URL$WEBHOOK_ID/$WEBHOOK_TOKEN"
    val payload = "{\"content\": \"$message\"}"

    val client = HttpClient()
    return client.post(url) {
        contentType(ContentType.Application.Json)
        setBody(payload)
    }
}
