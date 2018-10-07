from imutils import paths
import face_recognition
import argparse
import pickle
import cv2
import numpy as np
import os
import matplotlib.pyplot as plt

ap = argparse.ArgumentParser()
ap.add_argument("-p", "--picklefy", required=False, type=list)
ap.add_argument("-i", "--image", required=True)
ap.add_argument("-r", "--recognize", required=False,
	help="encoding to find (pickle file)",)
ap.add_argument("-d", "--detection-method", type=str, default="cnn",
	help="face detection model to use: either `hog` or `cnn`")
args = vars(ap.parse_args())

# Convert target hash on the chain from an array to a pickle 
if args["picklefy"]:
    p = np.array(args["picklefy"])
    with open("encodeing.pkl", "wb") as f:
        pickle.dump(p, f)
    exit(0)

print("[INFO] quantifying faces...") 
image = plt.imread(args["image"])

boxes = face_recognition.face_locations(image,
    model=args["detection_method"])

encoding = face_recognition.face_encodings(image, boxes)

# Compare current image to pickle file of target hash
if args["recognize"]:
    with open(args["recognize"], "rb") as f:
        target = pickle.load(f)
        #print(target)
    try:
        comparison = np.isclose(encoding[0], target)
        #print(comparison)
        print(np.all(comparison))
    except:
        print(False)

else:
    with open("encoding.pkl", "wb") as f:
        pickle.dump(encoding[0], f)