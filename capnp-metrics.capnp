using Go = import "/go.capnp";
$Go.package("metser");
$Go.import("github.com/jeromefroe/metser");

@0xc347c7e4235b72dc;

enum CapnpType {
  unknownType @0;
  counterType @1;
  batchTimerType @2;
  gaugeType @3;
}

struct CapnpCounter {
    id @0 :Data;
    value @1 :Int64;
}

struct CapnpBatchTimer {
    id @0 :Data;
    value @1 :List(Float64);
}

struct CapnpGauge {
    id @0 :Data;
    value @1 :Float64;
}

struct CapnpMetricUnion {
    type @0 :CapnpType;
    id @1 :Data;
    counterVal @2 :Int64;
    batchTimerVal @3 :List(Float64);
    gaugeVal @4 :Float64;
}

