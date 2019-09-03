import pandas
import numpy
import os
import matplotlib.pyplot as plt
from itertools import cycle
a = pandas.DataFrame([{"name": "shopspring\nDecimal", "num": 29},
{"name":"float32/float64" , "num": 26},
{"name":"int64" , "num": 15},
{"name":"big.Int", "num": 8 },
{"name":"crdb\napd.Decimal", "num": 7},
{"name":"big.Rat" , "num": 6},
{"name":"other"  , "num": 5},
{"name":"big.Float", "num": 4}])
x = numpy.arange(len(a.index))
# a = a.unstack()
a
plot = plt.bar(x, a.num)
plt.xticks(x, a.name)
plt.ylabel('number of people')
plt.title("Usage of different libraries for monetary purposes")
plt.xticks(rotation=65)
colors = cycle([(232/255, 62/255, 93/255, 0.3),(65/255, 186/255, 99/255, 0.3),(87/255, 126/255, 199/255, 0.3),(232/255, 165/255, 0, 0.3)])
for _, pl in enumerate(plot):
    pl.set_color(next(colors))
plt.savefig(os.path.abspath('../../content/')+f"/img/decimalSurvey.png", dpi=300, figsize=(50,25), bbox_inches='tight')