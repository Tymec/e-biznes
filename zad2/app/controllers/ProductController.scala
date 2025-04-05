package controllers

import models.Product

import javax.inject._
import play.api._
import play.api.mvc._
import play.api.libs.json._


@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
    private var products: Map[Long, Product] = Map()

    private def isValidProduct(obj: JsValue): Boolean = {
        val keys = obj.as[JsObject].keys
        val requiredKeys = Set("name", "description", "price")
        val missingKeys = requiredKeys.diff(keys)
        if (missingKeys.nonEmpty) {
            return false
        }
        return true
    }

    // GET /products
    def listProducts(): Action[AnyContent] = Action { implicit request =>
        Ok(Json.toJson(Json.obj("products" -> products.values.toList)))
    }

    // GET /products/:id
    def getProduct(id: Long): Action[AnyContent] = Action { implicit request =>
        products.get(id) match {
            case Some(product) => Ok(Json.toJson(product))
            case None => NotFound(Json.obj("error" -> "Product not found"))
        }
    }

    // POST /products
    def createProduct(): Action[JsValue] = Action(parse.json) { request =>
        isValidProduct(request.body) match {
            case true => {
                val product = request.body

                val name = (product \ "name").as[String]
                val description = (product \ "description").as[String]
                val price = (product \ "price").as[Double]
                
                products.values.find(_.name == name) match {
                    case Some(_) =>
                        BadRequest(Json.obj("error" -> "Product already exists"))
                    case None =>
                        val newId = products.size + 1L
                        val newProduct = Product(newId, name, description, price)
                        products += (newId -> newProduct)
                        Created(Json.toJson(newProduct))
                }
            }
            case false => BadRequest(Json.obj("error" -> "Invalid product data"))
        }
    }

    // PUT /products/:id
    def updateProduct(id: Long): Action[JsValue] = Action(parse.json) { request =>
        isValidProduct(request.body) match {
            case true => {
                val product = request.body

                val name = (product \ "name").as[String]
                val description = (product \ "description").as[String]
                val price = (product \ "price").as[Double]

                products.get(id) match {
                    case Some(_) =>
                        val updatedProduct = Product(id, name, description, price)
                        products += (id -> updatedProduct)
                        Ok(Json.toJson(updatedProduct))
                    case None =>
                        NotFound(Json.obj("error" -> "Product not found"))
                }
            }
            case false => BadRequest(Json.obj("error" -> "Invalid product data"))
        }
    }

    // DELETE /products/:id
    def deleteProduct(id: Long): Action[AnyContent] = Action { implicit request =>
        products.get(id) match {
            case Some(_) =>
                products -= id
                NoContent
            case None =>
                NotFound(Json.obj("error" -> "Product not found"))
        }
    }
}
