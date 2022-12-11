def main():
    with open('input.txt') as f:
        data = f.read().strip()
        chars = []
        for i, c in enumerate(data):
            chars.append(c)
            if len(chars) > 14:
                chars.pop(0)
            if len(chars) == 14 and len(set(chars)) == 14:
                print(i+1)
                break

if __name__ == '__main__':
    main()
