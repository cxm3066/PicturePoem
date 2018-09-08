import markovify

fh = open("OUTPUT.txt", "r")
lines = fh.readlines()

# Get raw text as string.
with open("poems.txt") as f:
    text = f.read()

# Build the model.
text_model = markovify.Text(text)

valid = False

while valid == False:
    valid = True
    out = text_model.make_short_sentence(70)
    for i in lines:
        if out in i:
            valid = False

print(out)