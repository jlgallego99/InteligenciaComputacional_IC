import urllib.request
import logging

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
	logging.info("Base de datos MNIST descargada en " + directorio)

def read_images():
	logging.info("Leyendo datos de MNIST...")

def main():
	descargar_mnist("./data/")

if __name__ == "__main__":
    main()