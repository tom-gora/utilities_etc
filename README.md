# utilities_etc
Various utilities helping me do the work.
Some I created, some repurposed or found.

## RIMG
> rimg

Stands for resize image. Created to speed up responsive design and varying size images creation for web. 
This bash script that will for each image passed as argument:

* Take image (images) or path (paths) to it (them) as argument(s).
* Create 3 copies scaled down with aspect ratio preserved to 0.75, 0.5 and 0.25 of the original size.
* Make a directory where the imgage being resized sits named like this image (for some_img.jpg an adjacent ./some_img/ will be made). If a directory of such name exists in the path it will exit and prompt you to either rid of it or rename the input file, to avoid risk of overwriting existing content.
* Move the original and resized output files in the newly created directory renaming them appropriately to indicate scaling factor.
* Prompt if operation successfull for each file.
 
