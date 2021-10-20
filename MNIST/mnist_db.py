import urllib.request
import logging

logging.basicConfig(level = logging.INFO)

MNIST_URL = "http://yann.lecun.com/exdb/mnist/"

# Datos de entrenamiento
trainingImages = "train-images-idx3-ubyte.gz"
trainingLabels = "train-labels-idx1-ubyte.gz"

# Datos de test
testImages = "t10k-images-idx3-ubyte.gz"
testLabels = "t10k-labels-idx1-ubyte.gz"
		
def descargar_mnist(directorio):
	logging.info("Descargando base de datos MNIST...")
	urllib.request.urlretrieve(MNIST_URL + trainingImages, directorio + trainingImages)
	urllib.request.urlretrieve(MNIST_URL + trainingLabels, directorio + trainingLabels)
	urllib.request.urlretrieve(MNIST_URL + testImages, directorio + testImages)
	urllib.request.urlretrieve(MNIST_URL + testLabels, directorio + testLabels)
	logging.info("Base de datos MNIST descargada en " + directorio)

def read_images():
	logging.info("Leyendo datos de MNIST...")

def main():
	descargar_mnist("./data/")

if __name__ == "__main__":
    main()