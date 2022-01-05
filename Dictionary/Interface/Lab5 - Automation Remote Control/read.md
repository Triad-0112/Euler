#Patch V 1.0.0
```
Problem :
  - Design the Home Automation Remote Control.
  - The remote consists of 7 programmable slots (each can be assigned to a different household device) along with corresponding on/off buttons for each.
  - Enclosed will be a set of APIs that is created by various vendors to control home automation devices such as lights, fans, audio equipments etc.
  - Note that it is important that we be able to control any future devices that the vendors may supply.
  - Seems like not a lot of thought has been given to come up with a set of common interfaces by the vendors.
  - We don’t want the remote to have to know the specifics of the vendor classes.
  - We don’t either want the remote to consist of a set of if statements, like “if slot1 == Light, then light.on(), else if slot1 = Hottub then    hottub.jetsOn()”

Solution Answer :
  - Decouple the requester of an action from the object that actually performs the action.
  - Here the requester would be the remote control and the object that performs the action would be an instance of one of the vendor class.
  - We can decouple them by introducing “command” objects into our design. A command object encapsulates a request to do something (like turn on a light) on a specific object (say, the living room light object). So, if we store a command object for each button, when the button is pressed we ask the command object to do some work.
  - The remote doesn’t have any idea what the work is, it just has a command object that knows how to talk to the right object to get the work done. So, you see, the remote is decoupled from the light object!
  - This pattern is called Command Pattern
  - Wiki Definition: The Command Pattern encapsulates a request as an object, thereby letting you parameterise other objects with different requests, queue or log requests, and support undoable operations.
```