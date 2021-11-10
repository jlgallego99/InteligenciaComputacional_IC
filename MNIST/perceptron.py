import tensorflow as tf
import keras
from keras.utils import to_categorical
import mnist_db as mnist

(training_images, training_labels, test_images, test_labels) = mnist.leer_mnist("./data/")

# Codificar las etiquetas
training_labels = to_categorical(training_labels)
test_labels = to_categorical(test_labels)