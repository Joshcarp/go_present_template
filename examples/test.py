import csv
import random
f = open("test.csv", "w")
write = csv.writer(f)

for i in range(10000):
    write.writerow([random.random()*random.random()**10, random.random()])
