import urllib.request
import logging
from glob import glob
import os
from mnist import MNIST

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
	
	archivos = glob(directorio + '*.gz')
	comando = 'gzip -dk ' + ' '.join(archivos)
	os.system(comando)

	logging.info("Base de datos MNIST descargada en " + directorio)
	
def leer_mnist(directorio):
	logging.info("Leyendo datos de MNIST...")
	datos_mnist = MNIST(directorio)

	# Cada imagen es una lista de bytes sin signo
	# Cada label es un array de bytes sin signo
	training_images, training_labels = datos_mnist.load_training()
	test_images, test_labels = datos_mnist.load_testing()

	logging.info("Datos de MNIST listos para usarse")

	# print(datos_mnist.display(training_images[0]))

def main():
	# descargar_mnist("./data/")
	leer_mnist("./data/")

if __name__ == "__main__":
    main()