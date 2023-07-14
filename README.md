# utilities_etc

Various utilities helping me do the work.
Some I created, some repurposed or found.

## RIMG

> rimg

Stands for resize image. Created to speed up responsive design and varying size images creation for web. It is intended to make prepping images for srcset atribute a breeze.

This bash script will for each image passed as argument do the following:

- Take image (images) or path (paths) to it (them) as argument(s).
- Create 3 extra copies scaled down with aspect ratio preserved to 0.75, 0.5 and 0.25 of the original size.
- Make a directory where the imgage being resized sits. The directory will be named like the image
  (for some_img.jpg an adjacent ./some_img/ will be made). If a directory of such name exists in the path, it will exit and prompt you to either rid of it or rename the input file, to avoid risk of overwriting existing content.
- Move the original and resized output files in the newly created directory renaming them appropriately to indicate scaling factor and the resulting resolution to quickly tell the size at the glance.
  <br><br>For instance:

  > some_img_25_270x180.jpg<br>
  > some_img_50_540x360.jpg<br>
  > some_img_75_810x540.jpg<br>
  > some_img_100_1080x720.jpg<br>

- Prompt if operation successfull for each file and prompt if errors occur.<br>
  (Errors checked for are: if arguments are missing, if passed files are images, if an output dir of given name exists already)

  ![demo_rimg](./assets/rimg.webp)
