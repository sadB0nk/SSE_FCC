#include <SPI.h>
#include "printf.h"
#include "RF24.h"

RF24 radio(7,8);

int data_get[3] = {0,0}; //data structure: [team number], [command number], [len of payload of elem massive], [payload massive]
int data_to_transmit[3] = {0,0};
byte address[][6] = {"1Node", "2Node", "3Node", "4Node", "5Node", "6Node"};
void setup() {
  Serial.begin(115200);
  printf_begin();
if (!radio.begin()){
  Serial.println("1Error of radio. Need to reconnect antenna and then restart arduino");
  while(1){}
}
  radio.setPALevel(RF24_PA_MAX);
  radio.setDataRate(RF24_1MBPS);
  radio.setChannel(0x70);
  for(int i=0; i<6; i++){
    radio.openReadingPipe(i, address[i]);
    Serial.print("1Was open reading pipe for team ");
    Serial.println(i+1);
  }
  radio.powerUp();
  radio.startListening();
  radio.setAutoAck(0);
  radio.printPrettyDetails();
}

void loop() {
  uint8_t pipe=0;
  
  if(Serial.available()){
    data_to_transmit[0] = Serial.parseInt();
    data_to_transmit[1] = Serial.parseInt();
    data_to_transmit[2] = Serial.parseInt();
    printf("1Was get data from computer: %d, %d\n",data_to_transmit[0], data_to_transmit[1]);
    radio.stopListening();
    radio.openWritingPipe(address[data_to_transmit[0]-1]);
    radio.write(data_to_transmit, sizeof(data_to_transmit));
    radio.openReadingPipe(0, address[0]);
    radio.startListening();    
  }

  if (radio.available(&pipe)){ 
    radio.read(data_get, sizeof(data_get));
    delay(20);
    Serial.print("1Was get message from team ");
    Serial.print(pipe+1);
    Serial.print(" with data: ");
    Serial.print(data_get[0]);
    Serial.print(" ");
    Serial.println(data_get[1]);
    delay(30);
    if (data_get[1] >= 1000){
      Serial.print("0Sending pong pocket to team ");
      Serial.print(pipe+1);
      Serial.print(" with data: ");
      Serial.println(data_get[1]-1000);
      radio.stopListening();
      radio.openWritingPipe(address[pipe]);
      data_to_transmit[0] = pipe+1;
      data_to_transmit[1] = data_get[1]-1000;
      data_to_transmit[2] = 0;
      radio.write(data_to_transmit, sizeof(data_to_transmit));
      radio.startListening();
    }
  }
  delay(5);
}
