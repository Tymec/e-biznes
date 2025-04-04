package models

import play.api.libs.json._


case class Product(
  id: Long,
  name: String,
  description: String,
  price: Double
)

object Product {
    implicit val productFormat: Format[Product] = Json.format[Product]
}
