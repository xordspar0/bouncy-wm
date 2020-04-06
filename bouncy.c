#include <X11/Xlib.h>
#include <stdbool.h>
#include <unistd.h>

struct direction {
    int x;
    int y;
};

int main()
{
    Display * display;
    Window root;
    Window parent;
    Window * windows;
    unsigned int nwindows;
    XWindowAttributes attr;
    int screenWidth, screenHeight;

    if(!(display = XOpenDisplay(0x0))) return 1;

    root = DefaultRootWindow(display);

    XGetWindowAttributes(display, root, &attr);
    screenWidth = attr.width;
    screenHeight = attr.height;

    XQueryTree(display, root, &root, &parent, &windows, &nwindows);

    struct direction directions[nwindows];
    for (int i = 0; i < nwindows; i++) {
        directions[i].x = 1;
        directions[i].y = 1;
    }

    for(;;) {
        for (int i = 0; i < nwindows; i++) {
            XGetWindowAttributes(display, windows[i], &attr);

            if (attr.x < 0 || attr.x + attr.width >= screenWidth) {
                directions[i].x *= -1;
            }
            if (attr.y < 0 || attr.y + attr.height >= screenHeight) {
                directions[i].y *= -1;
            }

            XMoveWindow(display, windows[i], attr.x + directions[i].x, attr.y + directions[i].y);
        }

        usleep(16700);
    }
}
