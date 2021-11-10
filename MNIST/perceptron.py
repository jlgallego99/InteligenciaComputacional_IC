import os
import logging
import tensorflow as tf
import tensorflow.keras
from tensorflow.keras.utils import to_categorical
import mnist_db as mnist

logging.basicConfig(level = logging.INFO)

# Descargar datos si no existen
if not os.path.isdir("./data/"):
    logging.info("El directorio de datos no existe, creando...")
    os.mkdir("./data/")
    mnist.descargar_mnist('./data/')
elif not os.listdir('./data/'):
    logging.info("El directorio de datos está vacío")
    mnist.descargar_mnist('./data/')
else:
    logging.info("Los datos ya están descargados")

(training_images, training_labels, test_images, test_labels) = mnist.leer_mnist("./data/")

# Codificar las etiquetas
training_labels = to_categorical(training_labels)
test_labels = to_categorical(test_labels)