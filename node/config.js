const fs = require('fs');

/**
 * Parallel Batch Processing in Node (Config)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
module.exports = {
    workersForImages: 19,
    workerFileImages: './workerImages.js',
    workerFileNumbers: './workerNumbers.js',
    imagePath: './images',
    slash: '/',
    imageResultPath: './imagesResult',
    files: fs.readdirSync("./images"),
    workerEventMessage: 'message',
    workerEventError: 'error',
    x: 800,
    y: 600,
    quality: 80,

    workersForNumbers: 10,
    numbers: Array(50000000).fill().map((x) => Math.floor(Math.random() * 10) + 1),
};