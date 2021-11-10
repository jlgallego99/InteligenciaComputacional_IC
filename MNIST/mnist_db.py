import urllib.request
import logging
from glob import glob
import os
from mnist import MNIST
import numpy as np

logging.basicConfig(level = logging.INFO)

MNIST_URL = "http://yann.lecun.com/exdb/mnist/"

# Datos de entrenamiento
training_images = "train-images-idx3-ubyte.gz"
training_labels = "train-labels-idx1-ubyte.gz"

# Datos de test
test_images = "t10k-images-idx3-ubyte.gz"
test_labels = "t10k-labels-idx1-ubyte.gz"
		
def descargar_mnist(directorio):
	logging.info("Descargando base de datos MNIST...")

	urllib.request.urlretrieve(MNIST_URL + training_images, directorio + training_images)
	urllib.request.urlretrieve(MNIST_URL + training_labels, directorio + training_labels)
	urllib.request.urlretrieve(MNIST_URL + test_images, directorio + test_images)
	urllib.request.urlretrieve(MNIST_URL + test_labels, directorio + test_labels)
	
	# Descomprimir datos
	archivos = glob(directorio + '*.gz')
	comando = 'gzip -dk ' + ' '.join(archivos)
	os.system(comando)

	# Eliminar comprimidos
	comando = 'rm ' + ' '.join(archivos)
	os.system(comando)

	logging.info("Base de datos MNIST descargada en " + directorio)
	
def leer_mnist(directorio):
	logging.info("Leyendo datos de MNIST...")
	datos_mnist = MNIST(directorio)

	# Cada imagen es una lista de bytes sin signo
	# Cada label es un array de bytes sin signo
	training_images, training_labels = datos_mnist.load_training()
	test_images, test_labels = datos_mnist.load_testing()

	# Convertir datos a arrays de numpy
	for i in range(len(training_images)):
		training_images[i] = np.reshape(training_images[i], (28, 28))

	for i in range(len(test_images)):
		test_images[i] = np.reshape(test_images[i], (28, 28))

	# Normalizar datos
	training_images = training_images.reshape((60000, 28 * 28))
	training_images = training_images.astype('float32') / 255
	test_images = test_images.reshape((10000, 28 * 28))
	test_images = test_images.astype('float32') / 255

	logging.info("Datos de MNIST listos para usarse")

	return (training_images, training_labels, test_images, test_labels)