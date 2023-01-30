package human.resource.mgmt.domain;

import human.resource.mgmt.domain.*;
import human.resource.mgmt.infra.AbstractEvent;
import java.util.*;
import lombok.*;

@Data
@ToString
public class EmployeeResigned extends AbstractEvent {

    private String userId;
    private String name;
    private String email;

    public EmployeeResigned(Employee aggregate) {
        super(aggregate);
    }

    public EmployeeResigned() {
        super();
    }
}
