
const { parentPort, workerData, isMainThread } = require("worker_threads");
const sharp = require('sharp');
const config = require('./config');

/**
 * Parallel Batch Processing in Node (Worker Script)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
function workerImages(images) {
  images.forEach(img =>
    sharp(`${config.imagePath}${config.slash}${img}`).resize(config.x, config.y).jpeg({ quality: config.quality })
        .toFile(`${config.imageResultPath}${config.slash}${img}`, (_, info) => {})
  );
  console.log("+++ WORKER FINISHED (processed images: %d)\n", images.length);
}

if (!isMainThread) { // verify call from worker thread
  if (!Array.isArray(workerData)) throw new Error("workerData must be an array of string"); // throw an error from worker
  parentPort.postMessage(workerImages(workerData));
}