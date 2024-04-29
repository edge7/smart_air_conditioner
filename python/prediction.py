import tensorflow as tf
import cv2
import numpy as np
model_reloaded = tf.keras.models.load_model("mobile_net_v3_small_rotation_minim.keras")

def get_prediction():
    img = cv2.cvtColor(cv2.imread("/tmp/fresh_image.jpg", cv2.IMREAD_UNCHANGED), cv2.COLOR_BGR2RGB)
    img = cv2.resize(img, (224, 224))
    pred = model_reloaded.predict(np.expand_dims(img, 0))[0]
    pred = round(float(pred), 2)
    print("prob is: ", pred)
    if pred > 0.5:
        pred = 'on'
    else:
        pred = 'off'
    print("pred is: ", pred)
    return pred