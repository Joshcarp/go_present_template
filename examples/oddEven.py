import numpy as np
import pandas
from collections import defaultdict
import matplotlib.pyplot as plt
import os

b = {}
for i in range(16770000, 16780000, 1000):
    b[i]={"odd":0, "even":0, "range":i}
    for j in range(i, i+1000):
        a = np.float32(j)
        print(a)
        if a%2 == 0:
            b[i]["even"] += 1
        else:
            b[i]["odd"] += 1
def no(s):
    a = s/1000
    a = a.astype(int)
    a = a.astype(str)
    return a + "e4"
df = pandas.DataFrame(b)
df = df.T
df["percent"] = df["even"]/(df["even"]+df["odd"])
p = df["percent"].plot()
plt.ylabel('Percentage')
plt.xlabel('range')
plt.title("Percentage of even numbers")
plt.ylim(top=1.1, bottom=0)
p.get_xaxis().get_major_formatter().set_useOffset(False)
p.get_xaxis().get_major_formatter().set_scientific(False)
plt.xticks(df["range"], no(df["range"]), rotation=50)
# plt.figure(figsize=(80,25))
plt.savefig(os.path.abspath('../content/')+f"/img/oddVsEven.png", dpi=300, figsize=(10,5),  bbox_inches = 'tight')
print(df.head())