from PIL import Image
import os

# 设置输入和输出目录
input_dir = './input'
output_dir = './out/'

# 循环输入目录中的所有JPEG图像
for filename in os.listdir(input_dir):

    # 只可以处理jpg和jpeg后缀下的程序，如果你的图片是别的后缀，请直接更改点后的文件格式
    if filename.endswith('.png') or filename.endswith('.jpg'):
        # 打开图像
        img = Image.open(os.path.join(input_dir, filename))
        mirrored_image = img.transpose(Image.FLIP_LEFT_RIGHT)
        mirrored_image.save(os.path.join(output_dir, filename))