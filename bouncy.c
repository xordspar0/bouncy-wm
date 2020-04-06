#include <X11/Xlib.h>
#include <stdbool.h>
#include <unistd.h>

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

    for(;;) {
        XGetWindowAttributes(display, root, &attr);
        screenWidth = attr.width;
        screenHeight = attr.height;

        XQueryTree(display, root, &root, &parent, &windows, &nwindows);
        for (int i = 0; i < nwindows; i++) {
            XGetWindowAttributes(display, windows[i], &attr);

            int x, y;
            if (attr.x + attr.width < screenWidth) {
                x = attr.x + 1;
            } else {
                x = attr.x - 1;
            }
            if (attr.y + attr.height < screenHeight) {
                y = attr.y + 1;
            } else {
                y = attr.y - 1;
            }

            XMoveWindow(display, windows[i], x, y);
        }

        usleep(16700);
    }
}
