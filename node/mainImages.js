const { Worker } = require('worker_threads');
const config = require('./config');

/**
 * Parallel Batch Processing in Node (mainImages)
 * @author Marco Guillén <mguillen.developer@gmail.com>
 */
async function mainImages() {
    function worker(images) {
        return new Promise((resolve, reject) => {
            const worker = new Worker(config.workerFileImages, { workerData: images });
            worker.on(config.workerEventMessage, resolve);
            worker.on(config.workerEventError, reject);
        });
    }

    const start = Date.now();
    const segmentsPerWorker = Math.ceil(config.files.length / config.workersForImages);
    const groupWorkers = Array(config.workersForImages).fill().map((_, index) => {
        let images;
        if (index === 0) {
            images = config.files.slice(0, segmentsPerWorker);
        } else if (index === config.workersForImages - 1) {
            images = config.files.slice(segmentsPerWorker * index);
        } else {
            images = config.files.slice(segmentsPerWorker * index, segmentsPerWorker * (index + 1))
        }
        return worker(images)
    });
    Promise.all(groupWorkers).then(_ => console.log(`::: INTENSIVE JOB SUCCESSFULLY COMPLETED -> Workers: ${config.workersForImages} -> Elapsed time: ${Date.now() - start} ms`)); // Wait for Workers and Promises Group to finish
}

mainImages();