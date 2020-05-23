#!/usr/bin/perl

my $COUNT = 0;
while (1) {
    print "Hello world $COUNT\n";

    $COUNT++;

	if ($COUNT > 10) {
		exit 1;
	}

    sleep 1;
}
