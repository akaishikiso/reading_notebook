#  value を直接変更する
class NumberBox:
    def __init__(self, value):
        self.value = value

# 引数として受け取ったポインタ変数に新しい構造体を代入（呼び出し元には影響しない）
def demonstrate_reference_behavior():
    box = NumberBox(10)
    update_value(box)
    print(box.value)  # 20　
    replace_box(box)
    print(box.value)  # 20（replaceBox の中で代入されたが元の number は変わっていない）
    empty_box = None
    replace_box(empty_box)
    print(empty_box is None)  # true（replaceBox 内で代入されても元の emptyBox は変わらない）

def update_value(box):
    box.value = 20

def replace_box(box):
    box = NumberBox(30)