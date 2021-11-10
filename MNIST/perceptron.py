import tensorflow as tf
import keras
import mnist_db as mnist

(training_images, training_labels, test_images, test_labels) = mnist.leer_mnist("./data/")