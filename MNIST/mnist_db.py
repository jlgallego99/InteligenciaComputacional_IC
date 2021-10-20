from os import urandom
import urllib.request
import logging
from glob import glob
import os

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
	
def main():
	descargar_mnist("./data/")
	# read_images("./data/")

if __name__ == "__main__":
    main()