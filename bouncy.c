#include <X11/Xlib.h>
#include <unistd.h>

int main()
{
    Display * dpy;
    Window root;
    Window * children;
    unsigned int nchildren;
    XWindowAttributes attr;

    if(!(dpy = XOpenDisplay(0x0))) return 1;

    root = DefaultRootWindow(dpy);

    for(;;)
    {
        XQueryTree(dpy, root, NULL, NULL, &children, &nchildren);
        usleep(167000);
    }
}
