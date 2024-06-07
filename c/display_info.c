#include <CoreGraphics/CoreGraphics.h>
#include <stdio.h>
#include "./display_info.h"

void printDisplayInfo(CGDirectDisplayID display) {
    CGSize size = CGDisplayScreenSize(display);
    CGRect bounds = CGDisplayBounds(display);
    printf("Display ID: %u\n", display);
    printf("Resolution: %.0fx%.0f\n", bounds.size.width, bounds.size.height);
    printf("Position: (%.0f, %.0f)\n", bounds.origin.x, bounds.origin.y);
    printf("Physical Size: %.0fx%.0f mm\n\n", size.width, size.height);
}

void listDisplays() {
    uint32_t displayCount;
    CGDirectDisplayID displays[32];
    
    CGGetActiveDisplayList(32, displays, &displayCount);
    
    printf("Number of active displays: %u\n\n", displayCount);
    
    for (uint32_t i = 0; i < displayCount; ++i) {
        printDisplayInfo(displays[i]);
    }
}

