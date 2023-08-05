a = 20


def func():
    global a
    a = 10
    print(a)


if __name__ == '__main__':
    func()
    print(a)
