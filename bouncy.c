#include <X11/Xlib.h>
#include <unistd.h>

int main()
{
    Display * display;
    Window root;
    Window parent;
    Window * windows;
    unsigned int nwindows;
    XWindowAttributes attr;

    if(!(display = XOpenDisplay(0x0))) return 1;

    root = DefaultRootWindow(display);

    for(;;) {
        XQueryTree(display, root, &root, &parent, &windows, &nwindows);
        for (int i = 0; i < nwindows; i++) {
            XGetWindowAttributes(display, windows[i], &attr);

            int x = attr.x + 1;
            int y = attr.y + 1;
            XMoveWindow(display, windows[i], x, y);
        }

        usleep(16700);
    }
}
