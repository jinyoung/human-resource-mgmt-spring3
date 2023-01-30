package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class EmployeeJoined extends AbstractEvent {

    private String userId;
    private String name;
    private String email;

    public EmployeeJoined(Employee aggregate) {
        super(aggregate);
    }

    public EmployeeJoined() {
        super();
    }
}
