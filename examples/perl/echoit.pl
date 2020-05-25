#!/usr/bin/perl
use Data::RandomPerson;
my $randomperson = Data::RandomPerson->new();

my $COUNT = 0;
while (1) {

    my $r = $randomperson->create();
    my $a = $r->{firstname};
    my $b = $r->{lastname};

    print "Hello world $COUNT $a $b\n";

    $COUNT++;

	if ($COUNT > 10) {
		exit 1;
	}

    sleep 1;
}
