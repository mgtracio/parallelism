package com.latamautos.challenge.benchmark

import java.io.File
import scala.util.Random

/**
 * Parallel Batch Processing in Scala (Config)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
object Config  {
  val WorkersForImages = 20
  val X = 800
  val Y = 600
  val Quality = 80
  val ImagePath = "images"
  val ImageResultPath = "imagesResult"
  val AbsolutePath = s"${new java.io.File(".").getCanonicalPath}/src/main/scala/com/latamautos/challenge/benchmark/$ImagePath"
  def imageWriter(img: File): File = new File(img.toString.replace(Config.ImagePath, Config.ImageResultPath))
  def images: List[File] = new File(AbsolutePath).listFiles.toList

  val WorkersForNumbers = 300
  def numbers: List[Int] = (1 to 30000000).map(_ => Random.between(1, 10)).toList
}